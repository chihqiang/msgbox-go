export interface AgentInfo {
  id: number;
  agent_no: string;
  agent_secret?: string;
  name?: string;
  phone?: string;
  email: string;
  status: boolean;
  created_at: string;
  updated_at: string;
}

export interface resetAgentSecret {
  agent_secret: string;
}
