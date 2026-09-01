import { defineConfig } from "orval";

export default defineConfig({
  api: {
    input: {
      target: "../openapi/openapi.yml",
      parserOptions: {
        externalRefs: {
          allow: ["*"],
        },
      },
    },
    output: {
      target: "./src/api/api.ts",
      schemas: "./src/api/model",
      client: "react-query",
      httpClient: "axios",
      baseUrl: "/api",
      override: {
        useBigInt: true,
      },
    },
  },
});
