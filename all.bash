#!/usr/bin/env bash

set -e

./make.bash

adb install -r bin/graaa-debug.apk

adb shell am start -a android.intent.action.MAIN \
    -n com.goandroidrpc.rpc/com.goandroidrpc.rpc.MainActivity
