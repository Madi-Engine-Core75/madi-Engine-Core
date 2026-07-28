use aes_gcm::{
    aead::{Aead, KeyInit, OsRng},
    Aes256Gcm, Nonce, Key // Or a fixed size type for key
};
use aes_gcm::aead::rand_core::RngCore;
use std::env;

pub struct SecureVault {
    cipher: Aes256Gcm,
}

impl SecureVault {
    /// تهيئة النواة الأمنية وجلب المفتاح من البيئة مع التحقق من صحة الطول
    pub fn new() -> Result<Self, String> {
        // جلب المفتاح السري من متغيرات البيئة بصرامة
        let key_hex = env::var("VAULT_SECRET_KEY")
            .map_err(|_| "CRITICAL: VAULT_SECRET_KEY environment variable is missing.")?;
        
        // التحقق من أن المفتاح يوافق متطلبات 256-bit (32 bytes)
        let key_bytes = hex::decode(key_hex)
            .map_err(|e| format!("Failed to decode hex key: {}", e))?;
            
        if key_bytes.len() != 32 {
            return Err("Invalid key length: AES-256 requires exactly 32 bytes.".into());
        }

        let key = aes_gcm::Key::<Aes256Gcm>::from_slice(&key_bytes);
        let cipher = Aes256Gcm::new(key);

        Ok(Self { cipher })
    }

    /// تشفير البيانات الحساسة مع توليد Nonce فريد لكل عملية لمنع هجمات التكرار
    pub fn encrypt(&self, plaintext: &[u8]) -> Result<(Vec<u8>, Vec<u8>), String> {
        let mut nonce_bytes = [0u8; 12];
        OsRng.fill_bytes(&mut nonce_bytes);
        let nonce = Nonce::from_slice(&nonce_bytes);

        let ciphertext = self.cipher.encrypt(nonce, plaintext)
            .map_err(|e| format!("Encryption failure: {}", e))?;

        Ok((ciphertext, nonce_bytes.to_vec()))
    }

    /// فك التشفير باستخدام نفس الـ Nonce والمفتاح
    pub fn decrypt(&self, ciphertext: &[u8], nonce_bytes: &[u8]) -> Result<Vec<u8>, String> {
        if nonce_bytes.len() != 12 {
            return Err("Invalid nonce length: must be 12 bytes.".into());
        }
        let nonce = Nonce::from_slice(nonce_bytes);

        let plaintext = self.cipher.decrypt(nonce, ciphertext)
            .map_err(|e| format!("Decryption failure or data tampering detected: {}", e))?;

        Ok(plaintext)
    }
}
