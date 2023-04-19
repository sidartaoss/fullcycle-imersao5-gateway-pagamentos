import { Box, Button, TextField, Typography } from "@mui/material";
import { FormEvent } from "react";
import axios from "axios";
import { useRouter } from "next/router";

const LoginPage = () => {
  const router = useRouter();
  async function onSubmitHandler(event: FormEvent) {
    event.preventDefault();

    const token = (document.querySelector("#token") as HTMLInputElement).value;

    try {
      await axios.post(`${process.env.NEXT_PUBLIC_API_HOST}/login`, { token });
      router.push("/orders");
    } catch (err) {
      console.log(err);
      alert("Login failed!");
    }
  }

  return (
    <Box
      sx={{
        marginTop: 8,
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Typography component="h1" variant="h5">
        Login
      </Typography>
      <Box component="form" onSubmit={onSubmitHandler} sx={{ mt: 1 }}>
        <TextField
          id="token"
          margin="normal"
          required
          fullWidth
          label="Token da conta"
        />
        <Button
          type="submit"
          fullWidth
          variant="contained"
          sx={{ mt: 3, mb: 2 }}
        >
          Login
        </Button>
      </Box>
    </Box>
  );
};

export default LoginPage;
