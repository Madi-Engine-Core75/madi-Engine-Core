use aes_gcm::{
    aead::{Aead, KeyInit, OsRng},
    Aes256Gcm, Nonce, Key
};
use aes_gcm::aead::rand_core::RngCore;

pub struct SecureVault;

impl SecureVault {
    /// تشفير البيانات الحساسة باستخدام AES-256-GCM
    pub fn encrypt_data(secret_key: &[u8; 32], plaintext: &[u8]) -> Result<(Vec<u8>, Vec<u8>), &'static str> {
        let key = Key::<Aes256Gcm>::from_slice(secret_key);
        let cipher = Aes256Gcm::new(key);
        
        let mut nonce_bytes = [0u8; 12];
        OsRng.fill_bytes(&mut nonce_bytes);
        let nonce = Nonce::from_slice(&nonce_bytes);

        let ciphertext = cipher.encrypt(nonce, plaintext)
            .map_err(|_| "Encryption failed")?;

        Ok((ciphertext, nonce_bytes.to_vec()))
    }
}
