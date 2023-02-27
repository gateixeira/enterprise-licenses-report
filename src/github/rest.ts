import {getOctokit} from '@actions/github'
import {token} from '../config/config'

export async function getLincensesInEnterprise(enterprise: string): Promise<unknown> {
  const octokit = getOctokit(token)

  const users = await octokit.request(`GET /enterprises/${enterprise}/consumed-licenses`)

  // eslint-disable-next-line no-console
  console.log(users)

  return users
}
