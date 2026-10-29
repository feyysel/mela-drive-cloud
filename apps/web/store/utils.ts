import { SerializedError } from "@reduxjs/toolkit";
import { FetchBaseQueryError } from "@reduxjs/toolkit/query";
import { ErrorResponse } from "./types";

export interface ApiErr extends ErrorResponse {
  status?: number;
}

function isFetchErr(err: unknown): err is FetchBaseQueryError {
  return typeof err === "object" && err !== null && "status" in err;
}

export function parseErr(
  err: FetchBaseQueryError | SerializedError | unknown,
): ApiErr {
  if (!err) {
    return { message: "An unexpected error occurred" };
  }

  if (isFetchErr(err)) {
    if (err.status === "FETCH_ERROR") {
      return {
        message: "Connection lost. Check your internet.",
        status: 0,
      };
    }

    if (err.status === "TIMEOUT_ERROR") {
      return {
        message: "The request timed out",
        status: 408,
      };
    }

    const ae = err.data as ErrorResponse | undefined;
    return {
      message: ae?.message || `Error ${err.status}`,
      status: typeof err.status === "number" ? err.status : undefined,
    };
  }

  const s = err as SerializedError;
  return {
    message: s.message || "An unexpected error occurred",
  };
}
