
<script lang="ts" context="module">
import { server } from "$app/env";


export function load({}) {
	if(!server && !IsLoggedIn){
		return {
			status: 302,
			redirect: "/",
		}
	}
}
</script>


<script lang="ts">
import { IsLoggedIn, Logout } from "$lib/user";

</script>

<button on:click={Logout}>Logout</button>
