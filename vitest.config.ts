import { defineConfig } from 'vitest/config';

export default defineConfig({
    test: {
        include: ['tests/**/*.test.ts'],
        environment: 'node',
        globals: true,
        passWithNoTests: true,
        reporters: 'verbose',
    }
})