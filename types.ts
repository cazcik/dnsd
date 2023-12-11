type Host = {
  name: string;
  ip: string;
};

type DomainData = {
  domain: string;
  host: Host[] | null;
  mx: Host[] | null;
  nameservers: Host[] | null;
  txt: string[];
};

export type ApiResponse =
  | {
      data: DomainData | undefined;
      message: string;
      status: string;
    }
  | undefined;
