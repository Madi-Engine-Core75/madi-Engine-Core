use aes_gcm::{
    aead::{Aead, KeyInit, OsRng},
    Aes256Gcm, Nonce, Key
};
use rand::RngCore;

pub struct CryptoEngine {
    key: Key<Aes256Gcm>,
}

impl CryptoEngine {
    pub fn new() -> Self {
        // توليد مفتاح تشفير عشوائي وآمن بطول 256-bit
        let key = Aes256Gcm::generate_key(&mut OsRng);
        Self { key }
    }

    pub fn encrypt(&self, data: &[u8]) -> Result<(Vec<u8>, Vec<u8>), &'static str> {
        let cipher = Aes256Gcm::new(&self.key);
        
        // توليد مفتاح عشوائي فريد لكل عملية تشفير (Nonce) بطول 12 بايت
        let mut nonce_bytes = [0u8; 12];
        OsRng.fill_bytes(&mut nonce_bytes);
        let nonce = Nonce::from_slice(&nonce_bytes);

        let ciphertext = cipher.encrypt(nonce, data)
            .map_err(|_| "Encryption failed")?;

        Ok((ciphertext, nonce_bytes.to_vec()))
    }

    pub fn decrypt(&self, ciphertext: &[u8], nonce_bytes: &[u8]) -> Result<Vec<u8>, &'static str> {
        let cipher = Aes256Gcm::new(&self.key);
        let nonce = Nonce::from_slice(nonce_bytes);

        let plaintext = cipher.decrypt(nonce, ciphertext)
            .map_err(|_| "Decryption failed")?;

        Ok(plaintext)
    }
}
