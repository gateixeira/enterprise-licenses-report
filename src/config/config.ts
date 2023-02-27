import {getInput} from '@actions/core'
import * as dotenv from 'dotenv'
dotenv.config()

export const enterprise: string = (process.env.ENTERPRISE as string) || getInput('enterprise', {required: true})
export const token: string = (process.env.PAT_TOKEN as string) || getInput('token', {required: true})
export const reportFormats: string[] = (
  (process.env.REPORT_FORMATS as string) || getInput('report-formats', {required: true})
).split(',')
