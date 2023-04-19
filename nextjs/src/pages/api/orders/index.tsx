import axios from "axios";
import { withIronSessionApiRoute } from "iron-session/next";
import { NextApiRequest, NextApiResponse } from "next";
import ironConfig from "../../../utils/iron-config";

export default withIronSessionApiRoute(ordersList, ironConfig);

async function ordersList(req: NextApiRequest, res: NextApiResponse) {
  const account = req.session.account;
  if (!account) {
    return res.status(401).json({ message: "Unauthenticated" });
  }
  try {
    const { data } = await axios.get(`${process.env.NEST_API_HOST}/orders`, {
      headers: {
        "x-token": account.token,
      },
    });
    res.status(200).json(data);
  } catch (err) {
    console.log(err);
    if (axios.isAxiosError(err)) {
      res.status(err.response!.status).json(err.response?.data);
    } else {
      res.status(500).json({ message: "Ocorreu um erro interno" });
    }
  }
}
