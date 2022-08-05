export const isValidBaritone = (data: Array<String>) => isValid(data)
// each command is an array of [command, maximum number of arguments]
const commands = {
    thisway: ["thisway", 1],
    goal: ["goal", 3],
    path: ["path", 0],
    cancel: ["cancel", 0],
    stop: ["stop", 0],
    mine: ["mine", 100],
    follow: ["follow", 2],
    wp: ["wp", 3],
    farm: ["farm", 2],
    explore: ["explore", 2],
    invert: ["invert", 0],
    come: ["come", 0],
    gc: ["gc", 0],
    find: ["find", 1]
}
const isValid = (command: Array<String>) => {
    let valid = false
    Object.entries(commands).forEach(cmd => {
        if (cmd[0] == command[0] && getArgs(command).length <= Number(cmd[1])) valid = true

    })
    return valid
}
// get the arguments of the command after the index 0
const getArgs = (command: Array<String>) => {
    return command.slice(1)
}