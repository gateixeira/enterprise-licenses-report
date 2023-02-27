import * as core from '@actions/core'
import {enterprise} from './config/config'
import {getLincensesInEnterprise} from './github/rest'
import {writeReport} from './report/summaryreport'

const run = async (): Promise<void> => {
  // get inputs
  core.debug(`[✅] Inputs parsed]`)

  core.info(`[✅] Getting license information for enterprise ${enterprise}`)

  const licenses = await getLincensesInEnterprise(enterprise)

  core.info(`[🔎] Found {licenses.length} licenses`)

  writeReport()

  core.info(`[✅] Summary Report written`)

  core.setOutput('report-json', JSON.stringify(licenses, null, 2))
  core.info(`[✅] Report written output 'report-json' variable`)

  return
}

run()
