<script>
import AlertStack from "$lib/components/alert-stack.svelte";
import Nav from "$lib/components/nav.svelte";
</script>

<Nav></Nav>

<div class="content">
	<slot></slot>
</div>
<AlertStack></AlertStack>

