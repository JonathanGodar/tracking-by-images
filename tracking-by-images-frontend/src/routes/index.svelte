<script lang="ts">
import { IsLoggedIn } from "$lib/user";
</script>
<h1>Index</h1>


{#if $IsLoggedIn}
<a href="/tracker/my-trackers">My trackers</a>
{/if}
