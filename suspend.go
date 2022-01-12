{{/* Suspend */}}
 
{{if eq (len .Args) 3}}
    {{ $arg1 := (index .CmdArgs 0) }}
    {{ $user := userArg $arg1 }}
	{{ $rolesremoved := cslice }}
 
    {{/* Staff Roles */}}
		{{/* Moderator */}}
			{{if targetHasRoleID $user.ID 487251866237927424}}
				{{ takeRoleID $user.ID 487251866237927424}}
				{{ $rolesremoved = $rolesremoved.Append "<@&487251866237927424>" }}
			{{end}}
		{{/* Recruiter */}}
			{{if targetHasRoleID $user.ID 520718391120625676}}
				{{ takeRoleID $user.ID 520718391120625676}}
				{{ $rolesremoved = $rolesremoved.Append "<@&520718391120625676>" }}
			{{end}}
		{{/* Staff */}}
			{{if targetHasRoleID $user.ID 667136206101807104}}
				{{ takeRoleID $user.ID 667136206101807104}}
				{{ $rolesremoved = $rolesremoved.Append "<@&667136206101807104>" }}
			{{end}}
			
	{{/* Moderator Roles */}}
		{{/* Game Moderator */}}
			{{if targetHasRoleID $user.ID 747898607570518056}}
				{{ takeRoleID $user.ID 747898607570518056}}
				{{ $rolesremoved = $rolesremoved.Append "<@&747898607570518056>" }}
			{{end}}
		{{/* Dead Matter Moderator */}}
			{{if targetHasRoleID $user.ID 714918684375777411}}
				{{ takeRoleID $user.ID 714918684375777411}}
				{{ $rolesremoved = $rolesremoved.Append "<@&714918684375777411>" }}
			{{end}}
		{{/* Deadside Moderator */}}
			{{if targetHasRoleID $user.ID 757706077876912178}}
				{{ takeRoleID $user.ID 757706077876912178}}
				{{ $rolesremoved = $rolesremoved.Append "<@&757706077876912178>" }}
			{{end}}
		{{/* EFT Moderator */}}
			{{if targetHasRoleID $user.ID 653670273240924171}}
				{{ takeRoleID $user.ID 653670273240924171}}
				{{ $rolesremoved = $rolesremoved.Append "<@&653670273240924171>" }}
			{{end}}
		{{/* EFT Instructor */}}
			{{if targetHasRoleID $user.ID 570325590847258664}}
				{{ takeRoleID $user.ID 570325590847258664}}
				{{ $rolesremoved = $rolesremoved.Append "<@&570325590847258664>" }}
			{{end}}
			
	{{/* Member Roles */}}
		{{/* Veteran Member */}}
			{{if targetHasRoleID $user.ID 518097300480983040}}
				{{ takeRoleID $user.ID 518097300480983040}}
				{{ $rolesremoved = $rolesremoved.Append "<@&518097300480983040>" }}
			{{end}}
		{{/* Member */}}
			{{if targetHasRoleID $user.ID 678362786115616769}}
				{{ takeRoleID $user.ID 678362786115616769}}
				{{ $rolesremoved = $rolesremoved.Append "<@&678362786115616769>" }}
			{{end}}
		{{/* Active Prospect */}}
			{{if targetHasRoleID $user.ID 678368898269052981}}
				{{ takeRoleID $user.ID 678368898269052981}}
				{{ $rolesremoved = $rolesremoved.Append "<@&678368898269052981>" }}
			{{end}}
		{{/* Prospect */}}
			{{if targetHasRoleID $user.ID 386341324372049932}}
				{{ takeRoleID $user.ID 386341324372049932}}
				{{ $rolesremoved = $rolesremoved.Append "<@&386341324372049932>" }}
			{{end}}
		{{/* Retired Member */}}
			{{if targetHasRoleID $user.ID 755104969778135200}}
				{{ takeRoleID $user.ID 755104969778135200}}
				{{ $rolesremoved = $rolesremoved.Append "<@&755104969778135200>" }}
			{{end}}
	
	{{/* Player Roles */}}
		{{/* EFT */}}
			{{if targetHasRoleID $user.ID 678368040278294542}}
				{{ takeRoleID $user.ID 678368040278294542}}
				{{ $rolesremoved = $rolesremoved.Append "<@&678368040278294542>" }}
			{{end}}
		{{/* Dead Matter */}}
			{{if targetHasRoleID $user.ID 678369475283648522}}
				{{ takeRoleID $user.ID 678369475283648522}}
				{{ $rolesremoved = $rolesremoved.Append "<@&678369475283648522>" }}
			{{end}}
		{{/* Hell Let Loose */}}
			{{if targetHasRoleID $user.ID 748296248838062221}}
				{{ takeRoleID $user.ID 748296248838062221}}
				{{ $rolesremoved = $rolesremoved.Append "<@&748296248838062221>" }}
			{{end}}
		{{/* Star Citizen */}}
			{{if targetHasRoleID $user.ID 698675838425497690}}
				{{ takeRoleID $user.ID 698675838425497690}}
				{{ $rolesremoved = $rolesremoved.Append "<@&698675838425497690>" }}
			{{end}}
		{{/* Deadside */}}
			{{if targetHasRoleID $user.ID 701096294650740848}}
				{{ takeRoleID $user.ID 701096294650740848}}
				{{ $rolesremoved = $rolesremoved.Append "<@&701096294650740848>" }}
			{{end}}
	
	{{/* Misc */}}
		{{/* [III%] */}}
			{{if targetHasRoleID $user.ID 745029683636207717}}
				{{ takeRoleID $user.ID 745029683636207717}}
				{{ $rolesremoved = $rolesremoved.Append "<@&745029683636207717>" }}
			{{end}}
		{{/* Streamer */}}
			{{if targetHasRoleID $user.ID 678369138237767728}}
				{{ takeRoleID $user.ID 678369138237767728}}
				{{ $rolesremoved = $rolesremoved.Append "<@&678369138237767728>" }}
			{{end}}
		{{/* Streaming Now */}}
			{{if targetHasRoleID $user.ID 667464135411367972}}
				{{ takeRoleID $user.ID 667464135411367972}}
				{{ $rolesremoved = $rolesremoved.Append "<@&667464135411367972>" }}
			{{end}}
		{{/* Giveaway Squad */}}
			{{if targetHasRoleID $user.ID 678369648911319073}}
				{{ takeRoleID $user.ID 678369648911319073}}
				{{ $rolesremoved = $rolesremoved.Append "<@&678369648911319073>" }}
			{{end}}
		{{/* Giveaways */}}
			{{if targetHasRoleID $user.ID 702636517880889384}}
				{{ takeRoleID $user.ID 702636517880889384}}
				{{ $rolesremoved = $rolesremoved.Append "<@&702636517880889384>" }}
			{{end}}
		{{/* Stream Alerts */}}
			{{if targetHasRoleID $user.ID 678369727587942424}}
				{{ takeRoleID $user.ID 678369727587942424 }}
				{{ $rolesremoved = $rolesremoved.Append "<@&678369727587942424>" }}
			{{end}}
		{{/* bots */}}
			{{if targetHasRoleID $user.ID 387336208344285184}}
				{{ takeRoleID $user.ID 387336208344285184 }}
				{{ $rolesremoved = $rolesremoved.Append "<@&387336208344285184>" }}
			{{end}}
 
    {{ giveRoleID $user.ID 678369126942638080 }}
	
    {{ $reason := (joinStr " " (slice .CmdArgs 1)) }}
    {{ $msg := (joinStr "" ($user.String) " has been suspended, all other roles have been removed. Roles will be reinstated after 24 hours.") }}
	{{ $rolesToReinstate := (joinStr "," $rolesremoved.StringSlice) }}
 
	{{ $channel := 678371934425055252 }}
    {{ $avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString .User.ID) "/" .User.Avatar ".png") }}
    {{ $sus_avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString $user.ID) "/" $user.Avatar ".png") }}
 
    {{ $embed := cembed "title" nil "description" nil "color" 0x5b76ff 
        "fields" (
            cslice
                ( sdict "name" "Suspension" "value" $msg )
                ( sdict "name" "Reason" "value" $reason )
				( sdict "name" "Roles Removed" "value" $rolesToReinstate )
        ) 
        "author" (sdict "name" (joinStr "" (.User.String) " (ID " (.User.ID) ")") "url" nil "icon_url" $avatar) "thumbnail" (sdict "url" $sus_avatar) "footer" nil "color" 16763904
    }}
 
    {{ sendMessage $channel $embed }}
	
	{{/* Trigger unsuspend after timer */}}
	{{ execCC 20 678371892750450701 86400 (sdict "user" $user "rolesRemoved" $rolesToReinstate "uname" $user.Username "udisc" $user.Discriminator) }}
{{ else }}
    {{ .User.Mention }}, please provide a reason in quotes!
{{ end }}
