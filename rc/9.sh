#!/bin/sh

PLAN9=${PLAN9:-$HOME/.local/plan9}

case ":$PATH:" in
	":$PLAN9/bin"*) ;;
	*) PATH=$PLAN9/bin:$PATH ;;
esac

export PATH PLAN9

