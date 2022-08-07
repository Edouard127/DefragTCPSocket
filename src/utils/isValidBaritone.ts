export const isValidBaritone = (data: Array<string>) => isValid(data)
// each command is an array of [command, maximum number of arguments]
const commands = {
    thisway: 1,
    goal: 3,
    goto: 3,
    path: 0,
    cancel: 0,
    stop: 0,
    mine: Infinity,
    follow: 2,
    wp: 3,
    farm: 2,
    explore: 2,
    invert: 0,
    come: 0,
    gc: 0,
    find: 1,
}
const isValid = (command: Array<string>) => {
    let valid = false

    const c = command[0]
    if (mapped().has(c)) {
        if(getArgs(command).length <= mapped().get(c)!) valid = true
    }

    return valid
}
const getArgs = (command: Array<string>) => {
    return command.slice(1)
}
const keys = Object.keys(commands)

const values = Object.values(commands)

const mapped = () => {
    const map = new Map<string, number>()
    keys.forEach((k, i) => map.set(k, values[i]))
    return map
}


