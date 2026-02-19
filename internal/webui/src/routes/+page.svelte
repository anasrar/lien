<script lang="ts">
	import { Button } from "$lib/components/ui/button";
	import { ArrowUpIcon, SearchIcon, SettingsIcon, SlashIcon, TrashIcon } from "@lucide/svelte";
	import {
		type ColumnDef,
		type ColumnFiltersState,
		getCoreRowModel,
		getFilteredRowModel,
		getSortedRowModel,
		type SortingState,
	} from "@tanstack/table-core";
	import { createSvelteTable, FlexRender, renderComponent } from "$lib/components/ui/data-table";
	import { TableRowEntryName } from "@/components/customs/table-row-entry-name";
	import { TableRowEntryDate } from "@/components/customs/table-row-entry-date";
	import { TableHeaderName } from "@/components/customs/table-header-name";
	import { TableHeaderModTime } from "@/components/customs/table-header-modtime";
	import { apiV1Client } from "@/openapi/clients/apiv1";
	import { createQuery } from "@tanstack/svelte-query";
	import type { components } from "@/openapi/schemas/apiv1";
	import { createSearchParam } from "$lib/hooks/search-param";
	import * as Breadcrumb from "$lib/components/ui/breadcrumb";
	import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
	import { setMode, mode } from "mode-watcher";
	import * as InputGroup from "$lib/components/ui/input-group/index.js";

	type DirectoryEntry = {
		isDirectory: boolean;
		modDate: Date;
	} & components["schemas"]["DirectoryEntry"];

	const [cwd, setCwd] = createSearchParam("cwd", "./");

	const breadcrumbs = $derived.by(() => {
		const tree = $cwd.slice(2).split("/");
		tree.pop();
		const result: { name: string; current: boolean; cwd: string }[] = [];

		const lastIndex = tree.length - 1;
		let i = 0;
		let treeStr = "./";
		for (const item of tree) {
			treeStr += item + "/";
			result.push({ name: item, current: i === lastIndex, cwd: treeStr });
			i++;
		}
		return result;
	});

	const query = createQuery(() => ({
		queryKey: ["ls", cwd],
		queryFn: async () => {
			return await apiV1Client.GET("/ls", {
				params: { query: { cwd: $cwd } },
			});
		},
	}));

	const columns: ColumnDef<DirectoryEntry>[] = [
		{
			accessorKey: "name",
			size: 200,
			header: ({ column }) => {
				const sortingHandler = column.getToggleSortingHandler()!;
				return renderComponent(TableHeaderName, {
					name: "Name",
					sortingHandler: sortingHandler,
					isSorted: column.getIsSorted(),
				});
			},
			cell: ({ row }) => {
				return renderComponent(TableRowEntryName, {
					isDirectory: row.original.isDirectory,
					name: row.original.name,
					cwd: $cwd,
					setCwd: setCwd,
				});
			},
		},
		{
			accessorKey: "modDate",
			header: ({ column }) => {
				const sortingHandler = column.getToggleSortingHandler()!;
				return renderComponent(TableHeaderModTime, {
					name: "Modification",
					sortingHandler: sortingHandler,
					isSorted: column.getIsSorted(),
				});
			},
			cell: ({ row }) => {
				return renderComponent(TableRowEntryDate, {
					modtime: row.original.modtime,
				});
			},
			sortingFn: "datetime",
		},
	];

	const entries = $derived.by(() => {
		const result: DirectoryEntry[] = [];
		if (query.isFetching) {
			return result;
		}
		if (query.data?.error !== undefined) {
			return result;
		}

		const { directories, files } = query.data!.data;

		result.push(
			...directories!.map<DirectoryEntry>((props) => {
				const modDate = new Date(props.modtime);
				return {
					isDirectory: true,
					modDate: modDate,
					...props,
					modtime: modDate.toISOString().slice(0, 19).replace("T", " "),
				};
			}),
			...files!.map<DirectoryEntry>((props) => {
				const modDate = new Date(props.modtime);
				return {
					isDirectory: false,
					modDate: modDate,
					...props,
					modtime: modDate.toISOString().slice(0, 19).replace("T", " "),
				};
			}),
		);

		return result;
	});

	let sorting = $state<SortingState>([]);
	let columnFilters = $state<ColumnFiltersState>([]);

	const table = createSvelteTable({
		get data() {
			return entries;
		},
		columns,
		getCoreRowModel: getCoreRowModel(),
		getSortedRowModel: getSortedRowModel(),
		getFilteredRowModel: getFilteredRowModel(),
		onSortingChange: (updater) => {
			if (typeof updater === "function") {
				sorting = updater(sorting);
			} else {
				sorting = updater;
			}
		},
		onColumnFiltersChange: (updater) => {
			if (typeof updater === "function") {
				columnFilters = updater(columnFilters);
			} else {
				columnFilters = updater;
			}
		},
		state: {
			get sorting() {
				return sorting;
			},

			get columnFilters() {
				return columnFilters;
			},
		},
	});
</script>

<div class="flex min-h-dvh flex-col font-mono">
	<div>
		<div class="mx-auto max-w-2xl">
			<div class="p-2">
				<div class="rounded-md border">
					<table class="w-full table-fixed">
						<thead class="sticky top-0">
							{#each table.getHeaderGroups() as headerGroup (headerGroup.id)}
								<tr>
									{#each headerGroup.headers as header (header.id)}
										<th
											colspan={header.colSpan}
											class={`${header.column.getIndex() === 1 ? "w-2/8 sm:w-1/3" : "w-6/8 sm:w-2/3"}`}
										>
											{#if !header.isPlaceholder}
												<FlexRender
													content={header.column.columnDef.header}
													context={header.getContext()}
												/>
											{/if}
										</th>
									{/each}
								</tr>
							{/each}
						</thead>
						<tbody class="divide-y">
							{#each table.getRowModel().rows as row (row.id)}
								<tr>
									{#each row.getVisibleCells() as cell (cell.id)}
										<td class="px-4 py-2">
											<FlexRender
												content={cell.column.columnDef.cell}
												context={cell.getContext()}
											/>
										</td>
									{/each}
								</tr>
							{:else}
								{#if query.isFetching}
									{#each Array(5)}
										<tr class="animate-pulse">
											<td class="px-4 py-2">
												<div class="h-6 w-32 rounded bg-gray-500"></div>
											</td>
											<td class="px-4 py-2">
												<div class="flex justify-center">
													<div class="h-6 w-6 sm:w-50 rounded bg-gray-500"></div>
												</div>
											</td>
										</tr>
									{/each}
								{:else}
									<tr>
										<td colspan={columns.length} class="h-24 text-center">
											{#if query.data?.error}
												{query.data?.error?.detail}
											{:else if entries.length === 0}
												Empty Folder
											{:else if table.getColumn("name")?.getFilterValue() !== ""}
												Search Not Found
											{/if}
										</td>
									</tr>
								{/if}
							{/each}
						</tbody>
					</table>
				</div>
			</div>
		</div>
	</div>
	<div class="sticky bottom-0 mt-auto border-t backdrop-blur-xl">
		<div class="mx-auto max-w-2xl p-2">
			<div class="flex flex-col gap-3">
				<div>
					<Breadcrumb.Root>
						<Breadcrumb.List>
							<Breadcrumb.Item>
								<button
									class="cursor-pointer"
									onclick={() => {
										if ($cwd === "./") {
											return;
										}
										setCwd("./");
									}}
								>
									root
								</button>
							</Breadcrumb.Item>
							<Breadcrumb.Separator>
								<SlashIcon />
							</Breadcrumb.Separator>
							{#each breadcrumbs as item}
								<Breadcrumb.Item>
									{#if item.current}
										<Breadcrumb.Page>{item.name}</Breadcrumb.Page>
									{:else}
										<button
											class="cursor-pointer"
											onclick={() => {
												if (item.current) {
													return;
												}

												setCwd(item.cwd);
											}}
										>
											{item.name}
										</button>
									{/if}
								</Breadcrumb.Item>
								{#if !item.current}
									<Breadcrumb.Separator>
										<SlashIcon />
									</Breadcrumb.Separator>
								{/if}
							{/each}
						</Breadcrumb.List>
					</Breadcrumb.Root>
				</div>
				<div>
					<InputGroup.Root>
						<InputGroup.Input
							placeholder="Filter..."
							value={(table.getColumn("name")?.getFilterValue() as string) ?? ""}
							onchange={(e) => {
								table.getColumn("name")?.setFilterValue(e.currentTarget.value);
							}}
							oninput={(e) => {
								table.getColumn("name")?.setFilterValue(e.currentTarget.value);
							}}
						/>
						<InputGroup.Addon>
							<SearchIcon />
						</InputGroup.Addon>
						{#if table.getColumn("name")?.getFilterValue()}
							<InputGroup.Addon align="inline-end">
								<InputGroup.Button
									variant="secondary"
									onclick={() => {
										table.getColumn("name")?.setFilterValue("");
									}}
								>
									<TrashIcon />
								</InputGroup.Button>
							</InputGroup.Addon>
						{/if}
					</InputGroup.Root>
				</div>
				<div class="flex flex-row justify-between gap-4">
					<div class="flex flex-row gap-2">
						<Button
							disabled={$cwd === "./"}
							size="icon"
							variant="outline"
							onclick={() => {
								setCwd("." + new URL("../", "http://example.com/" + $cwd).pathname);
							}}
						>
							<ArrowUpIcon />
						</Button>
					</div>
					<div class="flex flex-row gap-2">
						<DropdownMenu.Root>
							<DropdownMenu.Trigger>
								{#snippet child({ props })}
									<Button size="icon" variant="outline" {...props}><SettingsIcon /></Button>
								{/snippet}
							</DropdownMenu.Trigger>
							<DropdownMenu.Content class="w-32" align="end">
								<DropdownMenu.Group>
									<DropdownMenu.Sub>
										<DropdownMenu.SubTrigger>Theme</DropdownMenu.SubTrigger>
										<DropdownMenu.SubContent>
											<DropdownMenu.RadioGroup
												value={mode.current}
												onValueChange={(theme) => {
													setMode(theme as Parameters<typeof setMode>[0]);
												}}
											>
												<DropdownMenu.RadioItem value="light">Light</DropdownMenu.RadioItem>
												<DropdownMenu.RadioItem value="dark">Dark</DropdownMenu.RadioItem>
											</DropdownMenu.RadioGroup>
										</DropdownMenu.SubContent>
									</DropdownMenu.Sub>
								</DropdownMenu.Group>
							</DropdownMenu.Content>
						</DropdownMenu.Root>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
