<script lang="ts">
import type { AlertType } from "$lib/components/alert";

import AlertStack from "$lib/components/alert-stack.svelte";
import { writable } from "svelte/store";

const alerts = writable<AlertType[]>([]);

const handleAdd = () => {
	alerts.update((v) => [{
		message: toAdd, id: crypto.randomUUID()}, ...v])
}


let toAdd: string;
</script>

<input bind:value={toAdd} placeholder="alert type">
<button on:click={handleAdd}>Add Alert</button>


<AlertStack alerts={alerts}></AlertStack>
