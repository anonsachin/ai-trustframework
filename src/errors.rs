// The error values
#[derive(Debug)]
pub enum ExecutionError {
    InvalidArgError(String),
    ParseError(String),
    FileError(String),
}