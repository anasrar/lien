<script lang="ts" module>
	export type TableRowEntryDateProps = {
		modtime: string;
	};
</script>

<script lang="ts">
	import * as Tooltip from "$lib/components/ui/tooltip";
	import { buttonVariants } from "@/components/ui/button";
	import { ClockIcon } from "@lucide/svelte";

	let { modtime: date }: TableRowEntryDateProps = $props();
</script>

<div class="text-center">
	<span class="hidden whitespace-nowrap sm:inline">
		{date}
	</span>
	<Tooltip.Provider>
		<Tooltip.Root>
			<Tooltip.Trigger
				class={buttonVariants({ variant: "ghost", size: "icon-sm", class: "sm:hidden" })}
			>
				<ClockIcon />
			</Tooltip.Trigger>
			<Tooltip.Content>
				<p class="font-mono">{date}</p>
			</Tooltip.Content>
		</Tooltip.Root>
	</Tooltip.Provider>
</div>
