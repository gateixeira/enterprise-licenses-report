import * as core from '@actions/core'
import {enterprise} from './config/config'
import {addFileToRepo, getLicensesFile, getLincensesInEnterprise} from './github/rest'
import {writeReport} from './report/summaryreport'
import * as github from '@actions/github'
import {License} from './types/license'

const run = async (): Promise<void> => {
  // get inputs
  core.debug(`[✅] Inputs parsed]`)

  core.info(`[✅] Getting license information for enterprise ${enterprise}`)

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const licenses: any = await getLincensesInEnterprise(enterprise)

  const file: License[] = await getLicensesFile(github.context.repo.owner, github.context.repo.repo)

  core.info(`[🔎] Found {licenses.length} licenses`)

  file.concat({name: licenses.data.name})

  await addFileToRepo(github.context.repo.owner, github.context.repo.repo, JSON.stringify(file, null, 2))

  writeReport()

  core.info(`[✅] Summary Report written`)

  core.setOutput('report-json', JSON.stringify(licenses, null, 2))
  core.info(`[✅] Report written output 'report-json' variable`)

  return
}

run()
