<script lang="ts">

import { GetUserService } from "$lib/api";
import { IsLoggedIn, Login } from "$lib/user";

/* let loggedIn = IsLoggedIn(); */

let email: string;
let password: string;


const handleSignup = async () => {
	let response = await GetUserService().addUser({
			email, 
			password,
	});

	if(response.error) {
		alert(response.error)
		return;
	}

	let signinErr = await Login(email, password);
	if(signinErr != undefined){
		alert(signinErr.message);
	}
}
</script>

{#if $IsLoggedIn}
<p>You are already logged in. Please log out before creating a new account</p>
{:else}

<form on:submit|preventDefault={handleSignup}>

<input type="text" placeholder="email" bind:value={email}>
<input type="text" placeholder="password" bind:value={password}>

<button type="submit">Signup</button>
</form>
{/if}


<style>
form * {
  width: 100%;
  padding: 0.5rem 1.5rem;
  margin: 8px 0;
  box-sizing: border-box;

	border-color: black;
	border-style: solid;
	border-radius: 5rem; 

}
</style>

