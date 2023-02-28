import {getOctokit} from '@actions/github'
import {token} from '../config/config'
import {License} from '../types/license'

export async function getLincensesInEnterprise(enterprise: string): Promise<unknown> {
  const octokit = getOctokit(token)

  const users = await octokit.request(`GET /enterprises/${enterprise}/consumed-licenses`)

  // eslint-disable-next-line no-console
  console.log(users)

  return users
}

export async function getLicensesFile(owner: string, repository: string): Promise<License[]> {
  const octokit = getOctokit(token)

  const {data} = await octokit.rest.repos.getContent({
    owner,
    repo: repository,
    path: 'licenses.json'
  })

  // ignore non-directory responses
  if (!Array.isArray(data)) {
    return []
  }

  for (const item of data) {
    if (item.type === 'file') {
      // eslint-disable-next-line no-console
      console.log(item)

      return [{name: item.name}]
    }
  }

  return []
}

export async function addFileToRepo(owner: string, repository: string, content: string): Promise<void> {
  const octokit = getOctokit(token)

  const {data} = await octokit.rest.repos.getContent({
    owner,
    repo: repository,
    path: 'licenses.json'
  })

  // ignore non-directory responses
  if (!Array.isArray(data)) {
    return
  }

  //get date in format YYYY-MM-DD
  const date = new Date().toISOString().split('T')[0]
  const branch = `report-${date}`

  for (const item of data) {
    if (item.type === 'file') {
      await octokit.rest.repos.createOrUpdateFileContents({
        owner,
        repo: repository,
        path: 'licenses.json',
        message: `Update licenses`,
        branch,
        content: Buffer.from(JSON.stringify(content)).toString('base64'),
        sha: item.sha
      })
    }
  }

  //open pull request
  await octokit.rest.pulls.create({
    owner,
    repo: repository,
    title: 'Update licenses',
    head: branch,
    base: 'main'
  })
}
