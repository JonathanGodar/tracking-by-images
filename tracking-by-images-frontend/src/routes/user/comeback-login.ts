import { goto } from "$app/navigation";
import { page } from "$app/stores";
import { get } from "svelte/store";

export let originalPath = "";

export function ComebackLogin() {
	originalPath = get(page).url.pathname;
	goto("/user/login")
}


export async function Comeback(): Promise<boolean> {
	if(originalPath != ""){
		await goto(originalPath);
		originalPath = "";
		return true;
	}

	return false;
}
