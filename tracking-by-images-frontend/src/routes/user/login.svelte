<script lang="ts">
import { goto } from "$app/navigation";
import { session } from "$app/stores";

import { IsLoggedIn, Login } from "$lib/user";
import { Comeback, originalPath } from "./comeback-login";
let email: string, password: string;


export function onMount(){
	if($IsLoggedIn){
		goto("/")
		alert("You are already logged in")
	}
}
</script>

{#if $IsLoggedIn}
<p>You are already logged in. Please log out before creating a new account</p>
{:else}
<form on:submit|preventDefault={async () => {
	const err = await Login(email, password);
	if(err != undefined){
		alert(err.message);
		return;
	};

	let changed = await Comeback();
	if(!changed) {
			goto("/")
	}
}}>

<input type="email" placeholder="email" bind:value={email}>
<input type="password" placeholder="password" bind:value={password}>

<button type="submit">Login</button>
</form>
{/if}


