import { ReactComponent as AWSIcon } from 'images/chaos/aws.svg'
import { ReactComponent as ClockIcon } from 'images/chaos/time.svg'
import { ReactComponent as DNSIcon } from 'images/chaos/dns.svg'
import { ExperimentKind } from 'components/NewExperiment/types'
import { ReactComponent as FileSystemIOIcon } from 'images/chaos/io.svg'
import { ReactComponent as GCPIcon } from 'images/chaos/gcp.svg'
import { ReactComponent as LinuxKernelIcon } from 'images/chaos/kernel.svg'
import { ReactComponent as NetworkIcon } from 'images/chaos/network.svg'
import { ReactComponent as PodLifecycleIcon } from 'images/chaos/pod.svg'
import { ReactComponent as StressIcon } from 'images/chaos/stress.svg'
import { SvgIcon } from '@material-ui/core'
import T from 'components/T'

export function iconByKind(kind: ExperimentKind, size: 'small' | 'large' = 'large') {
  let icon

  switch (kind) {
    case 'PodChaos':
      icon = <PodLifecycleIcon />
      break
    case 'NetworkChaos':
      icon = <NetworkIcon />
      break
    case 'IoChaos':
      icon = <FileSystemIOIcon />
      break
    case 'KernelChaos':
      icon = <LinuxKernelIcon />
      break
    case 'TimeChaos':
      icon = <ClockIcon />
      break
    case 'StressChaos':
      icon = <StressIcon />
      break
    case 'DNSChaos':
      icon = <DNSIcon />
      break
    case 'AwsChaos':
      icon = <AWSIcon />
      break
    case 'GcpChaos':
      icon = <GCPIcon />
      break
  }

  return <SvgIcon fontSize={size}>{icon}</SvgIcon>
}

export function transByKind(kind: ExperimentKind) {
  switch (kind) {
    case 'PodChaos':
      return T('newE.target.pod.title')
    case 'NetworkChaos':
      return T('newE.target.network.title')
    case 'IoChaos':
      return T('newE.target.io.title')
    case 'KernelChaos':
      return T('newE.target.kernel.title')
    case 'TimeChaos':
      return T('newE.target.time.title')
    case 'StressChaos':
      return T('newE.target.stress.title')
    case 'DNSChaos':
      return T('newE.target.dns.title')
    case 'AwsChaos':
      return T('newE.target.aws.title')
    case 'GcpChaos':
      return T('newE.target.gcp.title')
  }
}
