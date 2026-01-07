import createClient from "openapi-fetch";
import type { paths } from "$lib/openapi/schemas/apiv1";

export const apiV1Client = createClient<paths>({ baseUrl: "/api/v1/" });
