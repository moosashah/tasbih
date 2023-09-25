// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

type Data = {
  message: string
}

const urlFrag = process.env['server'] || 'meep'

export default async function handler(
  _req: NextApiRequest,
  res: NextApiResponse<Data | unknown>
) {
  try {
    const url = `${urlFrag}/ping`
    const r = await fetch(url)
    const d: Data = await r.json()
    return res.status(200).json(d)
  } catch (e) {
    return res.status(500).json({ message: 'failed to ping backend', error: e })
  }
}
