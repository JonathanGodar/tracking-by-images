import { server } from "$app/env";
import { writable } from "svelte/store";
import { GetClient, GetUserService } from "./api";

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



export async function Login(email: string, password: string): Promise<Error | undefined> {
	try {
		let resp = await GetUserService().getAccessToken({
				email,
				password,
		});

		SetAccessToken(resp.token);
		return undefined;
	} catch(error: unknown){
		console.log(error);
		if(error instanceof Error) {
			return error;
		}
		throw error;
	}


}


function SetAccessToken(token: string){
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
