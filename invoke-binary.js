const childProcess = require('child_process')
const os = require('os')
const fs = require("fs")
const process = require('process')
const core = require('@actions/core');

function chooseBinary() {
    const platform = os.platform()
    const arch = os.arch()

    // otherwise assume it's a released version
    if (platform === 'linux' && arch === 'x64') {
        return `bin/status-writer-action_linux_amd64`
    }
    if (platform === 'linux' && arch === 'arm64') {
        return `bin/status-writer-action_linux_arm64`
    }
    if (platform === 'darwin' && arch === 'x64') {
        return `bin/status-writer-action_darwin_amd64`
    }
    if (platform === 'darwin' && arch === 'arm64') {
        return `bin/status-writer-action_darwin_arm64`
    }

    core.setFailed(`Unsupported platform (${platform}) and architecture (${arch})`)
}

function main() {
    const binary = chooseBinary()
    const mainScript = `${__dirname}/${binary}`
    const spawnSyncReturns = childProcess.spawnSync(mainScript, { stdio: 'inherit' })
    const status = spawnSyncReturns.status
    if (status !== 0) {
        core.setFailed(spawnSyncReturns.error)
    }
}

if (require.main === module) {
    main()
}
