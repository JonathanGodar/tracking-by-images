
<script lang="ts" context="module">
import { server } from "$app/env";
import { goto } from "$app/navigation";


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

<button on:click={() => {Logout(); goto("/user/login")}}>Logout</button>
