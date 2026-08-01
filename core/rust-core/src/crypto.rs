use aes_gcm::{
    aead::{Aead, KeyInit},
    Aes256Gcm, Nonce, Key
};
use rand::rngs::OsRng;
use rand::RngCore;

/// CryptoEngine provides AES-256-GCM encryption helpers for the core.
pub struct CryptoEngine {
    key: Key<Aes256Gcm>,
}

impl CryptoEngine {
    /// Initialize the engine with a pre-shared 32-byte key (from env or vault).
    pub fn new(secret_key: &[u8]) -> Result<Self, &'static str> {
        if secret_key.len() != 32 {
            return Err("Invalid key length: AES-256 requires exactly 32 bytes");
        }
        // Create a Key from the slice. `from_slice` returns a reference; dereference to obtain value.
        let key = *Key::<Aes256Gcm>::from_slice(secret_key);
        Ok(Self { key })
    }

    /// Generate a new random 32-byte master key (use for initial provisioning only).
    pub fn generate_master_key() -> Vec<u8> {
        let key = Aes256Gcm::generate_key(&mut OsRng);
        key.as_slice().to_vec()
    }

    /// Encrypt the provided plaintext bytes. Returns (ciphertext, nonce).
    pub fn encrypt(&self, data: &[u8]) -> Result<(Vec<u8>, Vec<u8>), &'static str> {
        let cipher = Aes256Gcm::new(&self.key);

        // Generate a unique 12-byte nonce for each encryption operation
        let mut nonce_bytes = [0u8; 12];
        OsRng.fill_bytes(&mut nonce_bytes);
        let nonce = Nonce::from_slice(&nonce_bytes);

        let ciphertext = cipher
            .encrypt(nonce, data)
            .map_err(|_| "Encryption failed")?;

        Ok((ciphertext, nonce_bytes.to_vec()))
    }

    /// Decrypt ciphertext using the provided nonce bytes.
    pub fn decrypt(&self, ciphertext: &[u8], nonce_bytes: &[u8]) -> Result<Vec<u8>, &'static str> {
        let cipher = Aes256Gcm::new(&self.key);
        let nonce = Nonce::from_slice(nonce_bytes);

        let plaintext = cipher
            .decrypt(nonce, ciphertext)
            .map_err(|_| "Decryption failed")?;

        Ok(plaintext)
    }
}
