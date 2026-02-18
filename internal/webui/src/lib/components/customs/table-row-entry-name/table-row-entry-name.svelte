<script lang="ts" module>
	export type TableRowEntryNameProps = {
		isDirectory: boolean;
		name: string;
		cwd: string;
		setCwd: (str: string) => void;
	};
</script>

<script lang="ts">
	import { FileIcon, FolderIcon } from "@lucide/svelte";

	let { isDirectory, name, cwd, setCwd }: TableRowEntryNameProps = $props();

	const ICON_SIZE = 16;

	const href = $derived.by(() => {
		return isDirectory
			? "./?cwd=" + encodeURIComponent(cwd + name + "/")
			: window.location.origin + "/dl" + (cwd + name).slice(1);
	});
</script>

<a
	{href}
	class="flex cursor-pointer flex-row items-center gap-2"
	onclick={(ev) => {
		ev.preventDefault();
		if (!isDirectory) {
			window.open(window.location.origin + "/dl" + (cwd + name).slice(1), "_blank");
			return;
		}
		setCwd(cwd + name + "/");
	}}
>
	{#if isDirectory}
		<FolderIcon size={ICON_SIZE} />
	{:else}
		<FileIcon size={ICON_SIZE} />
	{/if}
	{name}
</a>
