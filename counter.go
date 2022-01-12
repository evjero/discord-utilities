{{ $duration := toDuration .ExecData.Duration }}
{{ $msg := .ExecData.String }}
{{ $x := .ExecData.x }}
{{ $t := currentTime.Add ($duration) }}
{{ $mID := sendMessageRetID nil (joinStr  "" "Countdown starting..." $t.String) }}
{{ execCC 23 nil 0 (sdict "MessageID" $mID "T" $t "Message" ($msg) "x" ($x) ) }}
