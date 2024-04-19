Josh.ai | C++ Coding Challenge 2021
Below are the details for the challenge: 
The challenge is to create a program which implements a simple integration with the Philips Hue platform. It  should be a simple console based program that will print out text based on the state of the lights. When the  program starts up, it should print out all lights and their state (we will restrict to only whether the lights are  on/off, brightness, name, and ID). The output should be JSON in the format of the following  example: 
[
 {
 "name":"Red Lamp", 
 "id":"1", 
 "on":true, 
 "brightness":45 
 },
 {
 "name":"Green Lamp",
 "id":"2",
 "on":false,
 "brightness":100
 }
]
Your output does not need to be pretty printed like this, it can be a single line, however printing tab and newline  formatted JSON will probably make debugging easier for yourself. The “on” property is simply a boolean of  whether the light is on or not. The “brightness” property should be an integer (from 0 to 100) representing the  % brightness of the light. 
After printing out the initial state, your program should print out any changes in the lights’ states. For example, if  the Red Lamp turns off, you should print: 
{
 "id":"1",
 "on":false
}
If the Green Lamp is then turned on, and dimmed to 75%, you would print two changes: 
{
 "id":"2",
 "on":true
}
{
 "id":"2", 
 "brightness":75 
}

Your submission will be tested against the simulator found at https://www.npmjs.com/package/hue-simulator,  so please ensure that your solution works with it.  
We are purposefully not providing too much more information to you. Part of the challenge will be finding and  figuring out the documentation for the Philips Hue API since searching for this stuff is a real part of our job. Try to think through edge cases and how one would implement this in a real-world application.

Your program must be written in C++ (preferably C++14) and include all the necessary scripts, Makefiles and  instructions for us to build it. You should also include a ready-to-run, pre-built binary for some common  platform (Windows, macOS or Ubuntu/Debian). 
Note that even with modern versions of the standard library, you would probably have to pull other libraries in  to implement this. You are free to use any libraries you wish (for JSON, networking, etc.), unless they are  targeted specifically to the Philips Hue API. Please make sure to consider edge cases and think through how  the system will perform if various real-life issues arise. We are anticipating high quality code. 
Please send compiled binaries, the source code, and full instructions to compile and run the source  code. 

**A couple of additional notes on the challenge steps & time breakdown:  
Read and understand the challenge (in broad strokes) 
Download, setup and run the Hue simulator 
Search for and choose libraries to use for HTTP, JSON (and optionally command line parsing) (tip: look for "header-only" libraries) 
Implement the program 
Build and test the program 
Document the program 
What we are most interested in is how you tackle the actual problem, so we want you to spend most of  your time budget on your code, testing & documentation. 

For #2, the Hue simulator is a Node.js package. Here are some tips if you are unfamiliar with Node.js: 
To use on Windows: 
• Install Node.js: https://nodejs.org/en/download/ 
• Follow the instructions at: https://www.npmjs.com/package/hue-simulator 
To use on MacOS: 
• Install brew via: /usr/bin/ruby -e "$(curl - 
fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)" 
• Install Node.js: brew update && brew install node 
• Follow the instructions at: https://www.npmjs.com/package/hue-simulator 
To use on Linux: 
• Install Node.js: sudo apt update && sudo apt install nodejs && sudo apt install npm • Follow instructions at: https://www.npmjs.com/package/hue-simulator 
For #3, here are some suggestions for header-only libraries: 
• Simple C++ command line parser: https://github.com/FlorianRappl/CmdParser 
• C++11 header-only HTTP/HTTPS client library: https://github.com/yhirose/cpp-httplib • JSON for Modern C++: https://github.com/nlohmann/json 
It is a requirement that any libraries brought in are "header-only". The libraries do not have to use the ones  suggested here, but you must choose ones that are "header-only". This will keep the build process simple  and not waste time trying to figure out how to include, build or link static libraries.
Using the libraries we suggest  should reduce the time and guesswork involved in development, and should result in a concise solution of no more than a few hundred lines of code.
