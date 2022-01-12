{{if eq (len .Args) 2}}
	{{ $votedesc := (joinStr " " (slice .CmdArgs 0)) }}
	{{ $avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString .User.ID) "/" .User.Avatar ".png") }}
	{{ $embed := cembed "title" "NEW STAFF VOTE:" "description" $votedesc 
		"fields" (
			cslice
				( sdict "name" "Voting Options" "value" "In Favor = :arrow_up: \nAgainst = :arrow_down:\nIf voting against, please say why for re-evaluation" )
				( sdict "name" "Time Limit" "value" "48 Hours" )
		) 
		"author" (sdict "name" (joinStr "" (.User.String)) "url" nil "icon_url" $avatar) "thumbnail" (sdict "url" $avatar) "footer" nil "color" 0x5b76ff
		}}
 
	{{ $x := sendMessageRetID nil $embed }}
	{{addMessageReactions nil $x ":arrow_up:" ":arrow_down:"}}
	{{execCC 22 nil 0 (sdict "Duration" "48h" "String" "Time to vote!" "x" ($x)) }}	
{{ else }}
	{{ .User.Mention }}, please provide a vote title and description in two sets of quotes!
{{ end }}