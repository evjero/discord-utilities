{{/* Unsuspend */}}
{{ $user := .ExecData.user }}
 
{{/* Remove Suspended Role */}}
{{ if targetHasRoleID $user.ID 678369126942638080 }}
	{{ takeRoleID $user.ID 678369126942638080 }}
{{ end }}
 
{{/* Loop over removed roles and re-add them */}}
{{ range $index, $element := split .ExecData.rolesRemoved ","}}
	{{ $roleID := (slice $element 3 21) }}
	{{ giveRoleID $user.ID $roleID }}
{{ end }}
 
{{ $msg := (joinStr (.ExecData.uname) "> " " has been unsuspended") }}
{{ $avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString .User.ID) "/" .User.Avatar ".png") }}
{{ $sus_avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString $user.ID) "/" $user.Avatar ".png") }}
{{ $channel := 678371934425055252 }}
{{ $embed := cembed "title" nil "description" nil 
	"fields" (
		cslice
			( sdict "name" "Status" "value" $msg )
	) 
	"author" (sdict "name" (joinStr "" (.User.String) " (ID " (.User.ID) ")") "url" nil "icon_url" $avatar) "thumbnail" (sdict "url" $sus_avatar) "footer" nil "color" 16763904
}}
{{ sendMessage $channel $embed }}
