/*
 * Copyright (C) 2022 GalaxzyDev - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the AGPL-3.0 License.
 *
 * You should have received a copy of the AGPL-3.0 License with
 * this file. If not, please refer to the license linked below.
 * https://github.com/GalaxzyDev/Deadcord/blob/main/LICENSE.md
 */

package modules

import (
	"Deadcord/core"
	"Deadcord/requests"
	"Deadcord/util"
	"strconv"
)

func StartRemoveFriendThreads(user_id string) bool {
	wg.Add(len(core.RawTokensLoaded))
	for _, token := range core.RawTokensLoaded {
		go func(server_id string, token string) {
			defer wg.Done()
			removeFriendWorker(user_id, token)
		}(user_id, token)
	}

	wg.Wait()

	return false
}

func removeFriendWorker(user_id string, token string) {
	status, status_code, _ := requests.SendDiscordRequest("users/@me/relationships/"+user_id, "DELETE", token, "general", map[string]interface{}{}, map[string]interface{}{})

	if status {
		switch status_code {
		case 204:
			util.WriteToConsole("Bot successfully removed friend request.", 2)
		case 429:
			util.WriteToConsole("Removing friend request was limited.", 1)
		case 404:
			util.WriteToConsole("Could not find user to remove friend request.", 1)
		default:
			util.WriteToConsole("Bot could not remove friend request, request failed. Code: "+strconv.Itoa(status_code), 3)
		}
	}
}
