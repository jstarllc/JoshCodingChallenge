# Josh.ai | C++ Coding Challenge

## Challenge Details

The challenge is to create a program which implements a simple integration with a lights control platform (similar to Philips Hue). It should be a simple console-based program that will print out text based on the state of the lights. When the program starts up, it should print out all lights and their state. The output should be JSON in the format of the following example:
```json
[
    {
        "name": "Red Lamp", 
        "id": "1", 
        "room": "Living Room",
        "on": true, 
        "brightness": 45 
    },
    {
        "name": "Green Lamp",
        "id": "2",
        "room": "Office",
        "on": false,
        "brightness": 100
    }
]
```

Your output does not need to be pretty printed like this, it can be a single line, however printing tab and newline formatted JSON will probably make debugging easier for yourself. The `on` property is simply a boolean of whether the light is on or off. The `brightness` property in your output should be an integer (from 0 to 100) representing the % brightness of the light. Keep in mind that the brightness range used in the API is 0-255, so you will need to convert for your output.

After printing out the initial state, your program should print out any changes in the lights' states. For example, if the Red Lamp turns off, you should print:
```json
{
    "id": "1",
    "on": false
}
```

If the Green Lamp is then turned on, and dimmed to 75%, you would print two changes: 
```json
{
    "id": "2",
    "on": true
}
{
    "id": "2", 
    "brightness": 75 
}
```

Additionally, your program should detect when a light has been added or removed in the system. If a new light is discovered, print its full state:
```json
{
    "name": "Blue Lamp", 
    "id": "3", 
    "room": "Living Room",
    "on": true, 
    "brightness": 57 
}
```

If a light is removed, print a message containing the `id` and `name` of the light:
```
Blue Lamp (3) has been removed
```

**Your program should continue to monitor changes in the lights' states until stopped by the user.**

## Using the Simulator
Download the [latest release](https://github.com/jstarllc/JoshCodingChallenge/releases/latest) of the Lights Simulator. This is a server that maintains state for a collection of lights.

The Lights Simulator runs an HTTP server that you can send requests to interrogate and modify the state of the lights. This is representative of an IoT lighting system. Read through the [API Documentation](https://jstarllc.github.io/JoshCodingChallenge) to learn how to interract with the lights. 

There is a simple HTML page hosted by the Simulator. You can navigate to it in your browser at `localhost:8080` (or whatever IP and Port the Simulator is running on). This provides a simple GUI for you to interact with lights to test your application. Use this to trigger changes in the light state for your monitor program to print out.

## Deliverables
Your program must be written in C++ (preferably C++14) and include all the necessary scripts, Makefiles, and instructions for us to build it. You should also include a ready-to-run, pre-built binary for some common platform (Windows, macOS or Ubuntu/Debian).

Note that even with modern versions of the standard library, you will probably have to pull other libraries in to implement this. You are free to use any libraries you wish (for JSON, networking, etc.). **Any library you pull in must be header-only.** Please make sure to consider edge cases and think through how the system will perform if various real-life issues arise. We are anticipating high quality code.

**Please send compiled binaries, the source code, and full instructions to compile and run the source code.**

## Error Handling
Various errors come up when dealing with IoT devices over the network. Your application must be appropriately protected and not stop due to any such errors. Things to keep in mind:
* What happens when an HTTP request fails?
* What happens when JSON parsing fails?

## Challenge Steps and Time Breakdown
1. Read and understand the challenge (in broad strokes)
1. [Download](https://github.com/jstarllc/JoshCodingChallenge/releases/latest) and run the Simulator
1. Search for and choose header-only libraries to use for HTTP, JSON (and optionally command line parsing). Examples include:
    * [C++11 header-only HTTP/HTTPS client library](https://github.com/yhirose/cpp-httplib)
    * [JSON for Modern C++](https://github.com/nlohmann/json)
    * [Simple C++ command line parser](https://github.com/FlorianRappl/CmdParser)
1. Implement the program
1. Build and test the program
1. Document the program

**What we are most interested in is how you tackle the actual problem, so we want you to spend most of your time budget on your code, testing & documentation.**

**Using the libraries we suggest should reduce the time and guesswork involved in development, and should result in a concise solution of no more than a few hundred lines of code.**
