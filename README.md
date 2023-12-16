# Wisper Autohotkey Paste

Voice type anywhere in Windows using OpenAI's Whisper Speech Recognition engine!
This project allows you dictating anywhere in Windows using OpenAI's Whisper speech-to-text engine.

(This is a fork of mxro/autohotkey-chatgpt-voice, I modified it to allow just voice typing instead of carrying out commands.)

## Install

I have to update this section to match the current fork...

- Download and install AutoHotKey V1 from [autohotkey.com](https://www.autohotkey.com/)
- TBD
- Edit `config.json` from the extracted files. Provide your [Open API Key](https://www.howtogeek.com/885918/how-to-get-an-openai-api-key/) for the property `OpenapiKey`.

```json
{
  "OpenapiKey": "",
  "AutoHotKeyExec": ".\\bin\\autohotkey-1.1.37.01\\AutoHotkeyU64.exe"
}
```

## Usage

- Double click on `watch.ahk` from the extracted files
- Press F8
- Say whatever you want to type
- Press F8
- Wait for Open AI and AutoHotKey to do their magic

### Trigger Hotkey

The hotkey to start/stop a voice command is defined in `watch.ahk`. You can replace the following with a hotkey of your choice:

```
F8::
```

## Customise

- Edit transcriptionPrompt.txt to customize the transcription.
- Add multiple transcription prompt and .ahk files. For example, you can have multiple Trigger Hotkeys for different languages or contexts.

## Prior Art

- [ChatGPT-AutoHotkey-Utility](https://github.com/kdalanon/ChatGPT-AutoHotkey-Utility): Uses AutoHotKey to perform a number of actions, such as translate
- [ChatGPT Voice Assistant](https://github.com/DonGuillotine/chatGPT_whisper_AI_voice_assistant): Provides a Windows based assistant driven by ChatGPT
- [How to Make Your Own Windows Transcription App With Whisper and AutoHotkey](https://www.makeuseof.com/make-transcription-app-whisper-autohotkey/): Step by step tutorial to make a transcription app using AutoHotKey (added as per [reddit](https://www.reddit.com/r/AutoHotkey/comments/16ork8y/combining_ahk_with_chatgpt_to_automated_windows/))


## Develop

### Build Source Code

`task build`

### Package Executable

`task package`

### Run Locally

```
go run ./cmd/whisper-autohotkey/.
```
