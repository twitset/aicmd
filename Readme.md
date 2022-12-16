# aicmd 

Command line tool for talking to api.openai.com

Usage: aicmd [options] [command]

Options:
  -h, --help     display help for command

Example Commands:
    ./aicmd "What can you do on the command line?"
    or
    ./aicmd -max_tokens -temperature 0.6 "What can you do on the command line?"

Configuration:
    You can configure the model and other settings in the config.json file.