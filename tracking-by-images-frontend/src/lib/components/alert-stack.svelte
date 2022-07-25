<script lang="ts">
import { writable } from "svelte/store";
import type { AlertType } from "./alert";


import Alert from "./alert.svelte";

export let alerts = writable<AlertType[]>([]);

$: {
	
	if($alerts.length > 0){
		const id = ($alerts)[0].id
		setTimeout(() => {
			alerts.update(v => {
				return v.filter(v => v.id != id);
			});
			console.log("Removing stuff")
		}, 2500)
	}

}

</script>

<div class="stack">

{#each $alerts as alert (alert.id)}
	<Alert {...alert}></Alert>
{/each}
</div>


<style>
.stack {
	position: fixed; 

	display: flex;
	flex-direction: column-reverse;
	justify-content: space-between;


	/* debug */
	/* border: 2px dotted; */
	/* background-color: #FFFFFF22; */

	min-width: 20ch;

	bottom: 3vh;
	left: 5vh;

	z-index: 99999;
}

</style>
