<script lang="ts" module>
	import type { SortDirection } from "@tanstack/table-core";

	export type TableHeaderModTimeProps = {
		name: string;
		sortingHandler: (event: unknown) => void;
		isSorted: false | SortDirection;
	};
</script>

<script lang="ts">
	import { Button } from "@/components/ui/button";
	import { ArrowDown01Icon, ArrowDown10Icon, ClockIcon } from "@lucide/svelte";

	let { name, sortingHandler, isSorted }: TableHeaderModTimeProps = $props();
</script>

<div class="text-center">
	<Button variant="ghost" class="cursor-pointer text-center" onclick={sortingHandler}>
		<span class="hidden sm:inline">
			{name}
		</span>
		<ClockIcon class="inline sm:hidden" />
		{#if isSorted === "asc"}
			<ArrowDown01Icon />
		{:else if isSorted === "desc"}
			<ArrowDown10Icon />
		{/if}
	</Button>
</div>
