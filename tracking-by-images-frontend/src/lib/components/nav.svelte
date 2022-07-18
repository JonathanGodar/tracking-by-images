<script lang="ts">
import { page } from "$app/stores";
import { IsLoggedIn } from "$lib/user";
import { writable } from "svelte/store";



const paths = writable<{path: string, name: string}[]>([])

$: {
	paths.set([{
			path: "/",
			name: "Home",
	}])

	if($IsLoggedIn){
		paths.update((n) => {
				return [...n, {path: "/user/logout", name: "Logout"}];
			});
	} else {
		paths.update((n) => {
				return [...n, {path: "/user/login", name: "Login"}, {path: "/user/signup", name: "Signup"}];
		});
	}
	

	/* paths.push( */
	/* 	{ */
	/* 		path: "/", */
	/* 		name: "home", */
	/* 	} */
	/* ); */

	/* if($IsLoggedIn){ */
	/* 	paths.push({ */
	/* 		path: "/user/logout", */
	/* 		name: "Logout", */
	/* 	}); */
	/* } else { */
	/* 	paths.push( */
	/* 		{ */
	/* 			path: "/user/login", */
	/* 			name: "Login", */
	/* 		}, */
	/* 		{ */
	/* 			path: "/user/signup", */
	/* 			name: "Signup", */
	/* 		}); */
	/* } */
}


</script>


<nav>
	<div>
		<a class="tracker-site" href="/">TrackerSite</a>
		<ul>
			{#each $paths as path}
				<a class:active={path.path == $page.url.pathname} class="nav-item" href={path.path}>{path.name}</a>
			{/each}
		</ul>
	</div>
</nav>

<style>
* {
	background-color: #2C3333;
}

.tracker-site {
	color: #E7F6F2;
}

nav {
	padding: 1rem 2rem 0.7rem;
}

nav div {
	display: flex;
	justify-content: space-between;
}

nav a {
	margin-left: 0.2rem;
}

.tracker-site {
		text-decoration: bold;
}

.active {
	color: orange;
}

</style>
