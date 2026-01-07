<script lang="ts" module>
	import type { SortDirection } from "@tanstack/table-core";

	export type TableHeaderNameProps = {
		name: string;
		sortingHandler: (event: unknown) => void;
		isSorted: false | SortDirection;
	};
</script>

<script lang="ts">
	import { Button } from "@/components/ui/button";
	import { ArrowDownAZIcon, ArrowDownZAIcon } from "@lucide/svelte";

	let { name, sortingHandler, isSorted }: TableHeaderNameProps = $props();
</script>

<div>
	<Button variant="ghost" class="cursor-pointer text-center" onclick={sortingHandler}>
		{name}
		{#if isSorted === "asc"}
			<ArrowDownAZIcon />
		{:else if isSorted === "desc"}
			<ArrowDownZAIcon />
		{/if}
	</Button>
</div>
