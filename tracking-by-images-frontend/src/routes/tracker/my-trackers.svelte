<script lang="ts">
import { GetUserService } from "$lib/api";

import type { Tracker } from "$lib/oto-api";

import { IsLoggedIn } from "$lib/user";
import { get } from "svelte/store";


/* let trackers: Tracker[]; */
/* if(get(IsLoggedIn)) */

</script>

<h1>Your trackers!</h1>


{#if $IsLoggedIn}

{#await GetUserService().getMyTrackers() then resp}
	{resp.error}
	{resp.trackers}
	
{/await}
	
{/if}
