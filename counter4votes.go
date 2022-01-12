{{ $timeLeft := .ExecData.T.Sub currentTime }}
{{ $cntDownMessageHeader := joinStr "" "Countdown Timer: " .ExecData.Message }}
{{ $formattedTimeLeft := humanizeDurationSeconds $timeLeft }}

{{ $t := .ExecData.T }}
{{ $mID := .ExecData.MessageID }}
{{ $x := .ExecData.x }}
{{ $ts := .TimeSecond }}

{{ if lt $timeLeft (mult .TimeSecond 30) }}
  {{ range seq 1 (toInt $timeLeft.Seconds)  }}
    {{ $timeLeft := $t.Sub currentTime }}
    {{ $formattedTimeLeft := humanizeDurationSeconds $timeLeft }}

    {{ editMessage nil $mID (joinStr "" $cntDownMessageHeader "\nTime left: " $formattedTimeLeft) }}
    {{ if gt $timeLeft $ts }} {{ sleep 1 }} {{ end }}
  {{ end }}
  {{ editMessage nil  .ExecData.MessageID (joinStr "" $cntDownMessageHeader "\nTime left: **ENDED**") }}
  
  {{ $message := getMessage nil $x }}
	{{ $favor := 24 }}
	{{ $against := 0 }}
	{{ range $index, $element := $message.Reactions }}
		{{ if eq $index 0 }}
			{{ $favor = $element.Count }}
		{{ end }}
		{{ if eq $index 1 }}
			{{ $against = $element.Count }}
		{{ end }}
  {{ end }}
  {{ mentionRoleID 667136206101807104 }} Voting has concluded for ({{ $x }}):
	**In Favor:** {{ sub $favor 1 }}
  **Against:** {{ sub $against 1 }}
  
{{ else }}
    {{ editMessage nil .ExecData.MessageID (joinStr "" $cntDownMessageHeader "\nTime left: " $formattedTimeLeft) }}
    {{ execCC 23 nil 30 .ExecData }}
{{ end }}