// @ts-check

/**
 * @typedef {Object} Hasher
 * @property {(data: Uint8Array) => Promise<Uint8Array>} hash
 * @property {(hash: Uint8Array, data: Uint8Array) => Promise<boolean>} verify
 */
