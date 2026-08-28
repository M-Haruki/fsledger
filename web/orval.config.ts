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
      schemas: "./src/api/model.ts",
      client: "react-query",
      httpClient: "fetch",
      baseUrl: "/api",
    },
  },
});
