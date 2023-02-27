import * as core from '@actions/core'
//import {SummaryTableRow} from '@actions/core/lib/summary'

export function writeReport(): void {
  core.summary.addHeading('Summary')
  core.summary.addBreak()
  core.summary.addHeading('', 2)

  core.summary
    //.addHeading(name)
    //.addList(list)
    //.addHeading(heading, 2)
    // .addTable([
    //   tableHeaders.map(attribute => {
    //     return {data: attribute, header: true}
    //   }),
    //   ...tableBody
    // ] as SummaryTableRow[])
    .addBreak()
  core.summary.write()
}
