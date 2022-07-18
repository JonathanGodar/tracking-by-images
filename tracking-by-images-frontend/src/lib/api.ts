import { Client as apiClient, UserService as apiUserService, TrackerService as apiTrackerService} from "./oto-api"

// TODO Make to environment variable
export const endpoint = "http://localhost:8080/oto/"


let client: apiClient | undefined;

export function GetClient(): apiClient{
	if(client == undefined) {
		client = new apiClient()
		client.basepath = endpoint;
	}


	return client;
}


let userService: apiUserService | undefined;
export function GetUserService(): apiUserService {
	if(userService == undefined){
		userService = new apiUserService(GetClient());
	}

	return userService;
}

let trackerService: apiTrackerService | undefined;
export function GetTrackerService(): apiTrackerService {
	if(trackerService == undefined){
		trackerService = new apiTrackerService(GetClient());
	}

	return trackerService;
}
