#!/usr/bin/env python3

# Required parameters:
# @raycast.schemaVersion 1
# @raycast.title kebabify
# @raycast.mode silent

# Optional parameters:
# @raycast.icon 🐪
# @raycast.argument1 { "type": "text", "placeholder": "word" }
# @raycast.packageName lukasjoc

# Documentation:
# @raycast.description Turn CaMel to ke-bab case
# @raycast.author lukasjoc <jochamlu@gmail.com>
# @raycast.authorURL https://github.com/lukasjoc

import pyperclip
import re
import sys

result = re.sub(r'(?<!^)(?=[A-Z])', '-', sys.argv[1]).lower()
pyperclip.copy(result)
print(f"Copied {result} to clipboard")
