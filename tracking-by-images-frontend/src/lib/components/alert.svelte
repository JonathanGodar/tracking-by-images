<script lang="ts">
export let message: string;
export let type: "error" | "warning" | "info" = "info";
</script>


<div class="alert {type}">
{message}
</div>

<style>
.alert {
	margin: 20px 20px 0px 20px;
	/* margin-top: 20px; */
	padding: 0.4rem;

	display: flex;
	justify-content: center;

	font-size: 1.3rem;

	background-color: #00ffff22;
	border-radius: 25px;
}

.error {
	background-color: red;
}

.warning {
	background-color: yellow;
}

</style>


