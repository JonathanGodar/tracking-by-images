import { server } from "$app/env";
import { writable } from "svelte/store";
import { GetClient } from "./api";

const JWTStorageKey = "jwt";

export const IsLoggedIn = writable(GetIsLoggedIn())


function GetIsLoggedIn() {
	if (server) return false;

	const jwt = localStorage.getItem(JWTStorageKey)
	if(jwt != null){
		return true
	}

	return false;
}


export function SetAccessToken(token: string){
	localStorage.setItem(JWTStorageKey, token);


	GetClient().headers = (headers: Headers) => {
		headers.set("Authorization", `Bearer ${token}`)
	}

	IsLoggedIn.set(GetIsLoggedIn());
}

export function Logout(){
	localStorage.removeItem(JWTStorageKey);
	IsLoggedIn.set(GetIsLoggedIn());
}
