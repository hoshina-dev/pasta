-- Add processed_key column to part_3d_models
ALTER TABLE part_3d_models 
ADD COLUMN processed_key TEXT;

COMMENT ON COLUMN part_3d_models.processed_key IS 'S3 key for the optimized/processed 3D model file';

-- Create optimization_job_logs table for BI and monitoring
CREATE TABLE IF NOT EXISTS optimization_job_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    job_id UUID NOT NULL,
    
    -- Job configuration
    source_url TEXT NOT NULL,
    dest_url TEXT NOT NULL,
    source_key TEXT,
    dest_key TEXT,
    
    -- Optimization parameters
    draco_compression_level INT,
    draco_position_quantization INT,
    draco_texcoord_quantization INT,
    draco_normal_quantization INT,
    draco_generic_quantization INT,
    
    -- Job execution details
    status TEXT NOT NULL,
    exit_code INT,
    error_message TEXT,
    
    -- File metrics
    source_file_size BIGINT,
    processed_file_size BIGINT,
    compression_ratio DECIMAL(5,2),
    
    -- Timing metrics
    started_at TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    duration_seconds INT,
    
    -- Complete logs for debugging
    job_logs TEXT,
    
    -- Metadata
    webhook_received_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for BI queries
CREATE INDEX IF NOT EXISTS idx_optimization_job_logs_job_id ON optimization_job_logs(job_id);
CREATE INDEX IF NOT EXISTS idx_optimization_job_logs_status ON optimization_job_logs(status);
CREATE INDEX IF NOT EXISTS idx_optimization_job_logs_created_at ON optimization_job_logs(created_at);
CREATE INDEX IF NOT EXISTS idx_optimization_job_logs_completed_at ON optimization_job_logs(completed_at);
CREATE INDEX IF NOT EXISTS idx_optimization_job_logs_duration ON optimization_job_logs(duration_seconds);

-- Comments for documentation
COMMENT ON TABLE optimization_job_logs IS 'Complete log of all 3D optimization jobs for BI analysis and monitoring';
COMMENT ON COLUMN optimization_job_logs.job_id IS 'UUID of the optimization job (matches part_3d_models.id)';
COMMENT ON COLUMN optimization_job_logs.status IS 'Job status: success, failed';
COMMENT ON COLUMN optimization_job_logs.exit_code IS 'Process exit code from Argo job';
COMMENT ON COLUMN optimization_job_logs.compression_ratio IS 'Percentage reduction in file size (0-100)';
COMMENT ON COLUMN optimization_job_logs.duration_seconds IS 'Total job execution time in seconds';
COMMENT ON COLUMN optimization_job_logs.job_logs IS 'Complete stdout/stderr logs from Argo job execution';
