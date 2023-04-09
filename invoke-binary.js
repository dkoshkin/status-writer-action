const childProcess = require('child_process')
const os = require('os')
const process = require('process')

const VERSION = 'v0.1.0'

function chooseBinary() {
    const platform = os.platform()
    const arch = os.arch()

    if (platform === 'linux' && arch === 'x64') {
        return `gha-monitor_linux_amd64_v1/gha-monitor`
    }
    if (platform === 'linux' && arch === 'arm64') {
        return `gha-monitor_linux_arm64/gha-monitor`
    }
    if (platform === 'darwin' && arch === 'x64') {
        return `gha-monitor_darwin_adm64_v1/gha-monitor`
    }
    if (platform === 'darwin' && arch === 'arm64') {
        return `gha-monitor_darwin_arm64/gha-monitor`
    }
    if (platform === 'windows' && arch === 'x64') {
        return `gha-monitor_windows_adm64_v1/gha-monitor.exe`
    }
    if (platform === 'windows' && arch === 'arm64') {
        return `gha-monitor_windows_arm64/gha-monitor.exe`
    }

    console.error(`Unsupported platform (${platform}) and architecture (${arch})`)
    process.exit(1)
}

function main() {
    const binary = chooseBinary()
    const mainScript = `${__dirname}/dist/${binary}`
    const spawnSyncReturns = childProcess.spawnSync(mainScript, { stdio: 'inherit' })
    const status = spawnSyncReturns.status
    if (typeof status === 'number') {
        process.exit(status)
    }
    process.exit(1)
}

if (require.main === module) {
    main()
}
