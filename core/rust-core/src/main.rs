use tonic::{transport::Server, Request, Response, Status};

// استيراد الملفات المولدة تلقائياً من الـ proto
pub mod madi {
    pub mod engine {
        pub mod v1 {
            tonic::include_proto!("madi.engine.v1");
        }
    }
}

use madi::engine::v1::{
    madi_engine_core_server::{MadiEngineCore, MadiEngineCoreServer},
    PayloadRequest, PayloadResponse, HealthRequest, HealthResponse,
};

#[derive(Default)]
pub struct EngineService {}

#[tonic::async_trait]
impl MadiEngineCore for EngineService {
    async fn process_payload(
        &self,
        request: Request<PayloadRequest>,
    ) -> Result<Response<PayloadResponse>, Status> {
        let req = request.into_inner();
        println!("Received request_id: {} with action: {}", req.request_id, req.action_type);

        // معالجة البيانات (تجهيزاً لدمج تشفير AES-256-GCM لاحقاً)
        let processed_payload = req.encrypted_data;

        let reply = PayloadResponse {
            request_id: req.request_id,
            success: true,
            processed_data: processed_payload,
            error_message: "".to_string(),
        };

        Ok(Response::new(reply))
    }

    async fn health_check(
        &self,
        _request: Request<HealthRequest>,
    ) -> Result<Response<HealthResponse>, Status> {
        let reply = HealthResponse {
            status: "ONLINE - MadiEngineCore is operational".to_string(),
            timestamp: chrono::Utc::now().timestamp(),
        };
        Ok(Response::new(reply))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse()?;
    let service = EngineService::default();

    println!("MadiEngineCore gRPC server listening on {}", addr);

    Server::builder()
        .add_service(MadiEngineCoreServer::new(service))
        .serve(addr)
        .await?;

    Ok(())
}
